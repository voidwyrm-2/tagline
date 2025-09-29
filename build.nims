#!/usr/bin/env nim --hints:off

import std/[
  strutils,
  strformat
]

type
  Os = enum
    oLinux,
    oDarwin,
    oWindows,
    oFreebsd,
    oOpenbsd,
    oNetbsd,
    oWasip1

  Arch = enum
    a386,
    aAmd64,
    aArm,
    aArm64,
    aRiscv64,
    aWasm

template f(v: untyped): untyped =
  ($v)[1..^1].toLower

let
  appname = "tagline"
  rmCmd = fmt"rm {appname}"
  targets = [
    (oLinux, @[a386, aAmd64, aArm]),
    (oFreebsd, @[a386, aAmd64, aArm]),
    (oOpenbsd, @[a386, aAmd64, aArm]),
    (oNetbsd, @[a386, aAmd64, aArm]),
    (oDarwin, @[aArm64]),
    (oWindows, @[aAmd64])
  ]

exec "rm -rf out/"
mkDir "out"

for (os, arches) in targets:
  for arch in arches:
    echo "Building ", f(os), "/", f(arch), "..."
    exec fmt"GOOS={f(os)} GOARCH={f(arch)} go build -o out/{appname} ."
    echo "Built ", f(os), "/", f(arch), ", zipping..."
    withDir "out/":
      exec fmt"zip {f(os)}-{f(arch)} {appname}"
      echo "Zipped"
      exec rmCmd
