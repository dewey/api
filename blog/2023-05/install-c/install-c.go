package main

import (
   "2a.pages.dev/nursery"
   "fmt"
   "github.com/klauspost/compress/zstd"
   "io"
   "net/http"
   "os"
   "path/filepath"
)

func extract(in, out string) error {
   file, err := os.Open(in)
   if err != nil {
      return err
   }
   defer file.Close()
   fmt.Println("Zstd", in)
   read, err := zstd.NewReader(file)
   if err != nil {
      return err
   }
   return nursery.Tar(read, out, 1)
}

func download(in, out string) error {
   fmt.Println("Stat", out)
   if _, err := os.Stat(out); err == nil {
      return nil
   }
   fmt.Println("Get", in)
   res, err := http.Get(in)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return os.WriteFile(out, data, 0666)
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.Join(home, "nursery/c")
   for _, file := range files {
      home_file := filepath.Join(home, file)
      err := download(mirror + file, home_file)
      if err != nil {
         panic(err)
      }
      if err := extract(home_file, `D:\c`); err != nil {
         panic(err)
      }
   }
}

// avoid redirect, also clarkson mirror sucks
const mirror = "http://mirror.umd.edu/msys2/mingw/mingw64/"

var files = []string{
   // fatal error: cannot find 'ld'
   "mingw-w64-x86_64-binutils-2.38-2-any.pkg.tar.zst",
   // cannot find crt2.o: No such file or directory
   "mingw-w64-x86_64-crt-git-9.0.0.6451.a3f6d363d-1-any.pkg.tar.zst",
   // linker `x86_64-w64-mingw32-gcc` not found
   "mingw-w64-x86_64-gcc-11.2.0-10-any.pkg.tar.zst",
   // libgmp-10.dll was not found
   "mingw-w64-x86_64-gmp-6.2.1-3-any.pkg.tar.zst",
   // fatal error: stddef.h: No such file or directory
   "mingw-w64-x86_64-headers-git-9.0.0.6451.a3f6d363d-1-any.pkg.tar.zst",
   // libwinpthread-1.dll was not found
   "mingw-w64-x86_64-libwinpthread-git-9.0.0.6451.a3f6d363d-1-any.pkg.tar.zst",
   // cannot find default-manifest.o: No such file or directory
   "mingw-w64-x86_64-windows-default-manifest-6.4-4-any.pkg.tar.zst",
   // cannot find -l:libpthread.a: No such file or directory
   "mingw-w64-x86_64-winpthreads-git-9.0.0.6451.a3f6d363d-1-any.pkg.tar.zst",
   // zlib1.dll was not found
   "mingw-w64-x86_64-zlib-1.2.11-9-any.pkg.tar.zst",
   // libzstd.dll was not found
   "mingw-w64-x86_64-zstd-1.5.2-2-any.pkg.tar.zst",
}
