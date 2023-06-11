package nursery

import (
   "archive/tar"
   "archive/zip"
   "github.com/xi2/xz"
   "io"
   "os"
   "path/filepath"
)

func create(head *tar.Header, in io.Reader, out string) error {
   head.Name = filepath.Join(out, head.Name)
   err := os.MkdirAll(filepath.Dir(head.Name), 0666)
   if err != nil {
      return err
   }
   switch head.Typeflag {
   case tar.TypeReg:
      data, err := io.ReadAll(in)
      if err != nil {
         return err
      }
      return os.WriteFile(head.Name, data, 0777)
   case tar.TypeLink:
      _, err := os.Stat(head.Name)
      if err != nil {
         return os.Link(filepath.Join(out, head.Linkname), head.Name)
      }
   }
   return nil
}

func Xz(in, dir string, level int) error {
   file, err := os.Open(in)
   if err != nil {
      return err
   }
   defer file.Close()
   read, err := xz.NewReader(file, 0)
   if err != nil {
      return err
   }
   return Tar(read, dir, level)
}

func strip(in string, n int) string {
   for i, char := range in {
      if n <= 0 {
         return in[i:]
      }
      if char == '/' {
         n--
      }
   }
   return ""
}

// Need for Zstandard
func Tar(in io.Reader, dir string, level int) error {
   read := tar.NewReader(in)
   for {
      head, err := read.Next()
      if err == io.EOF {
         break
      } else if err != nil {
         return err
      }
      head.Name = strip(head.Name, level)
      if head.Name != "" {
         head.Linkname = strip(head.Linkname, level)
         err := create(head, read, dir)
         if err != nil {
            return err
         }
      }
   }
   return nil
}

func Zip(in, dir string, level int) error {
   read, err := zip.OpenReader(in)
   if err != nil {
      return err
   }
   defer read.Close()
   for _, head := range read.File {
      if head.Mode().IsDir() {
         continue
      }
      rc, err := head.Open()
      if err != nil {
         return err
      }
      data, err := io.ReadAll(rc)
      if err != nil {
         return err
      }
      if err := rc.Close(); err != nil {
         return err
      }
      head.Name = filepath.Join(dir, strip(head.Name, level))
      if err := os.MkdirAll(filepath.Dir(head.Name), 0666); err != nil {
         return err
      }
      if err := os.WriteFile(head.Name, data, 0777); err != nil {
         return err
      }
   }
   return nil
}
