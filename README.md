# go-vigem

Go bindings for the ViGEmClient library.

## Requirements

This library requires the [ViGEmClient](https://github.com/ViGEm/ViGEmClient) and [ViGEmBus](https://github.com/ViGEm/ViGEmBus) to work.

- Download and install [ViGEmBus](https://github.com/ViGEm/ViGEmBus/releases)
- Download ViGEmClient.dll
    - [x64](https://ci.appveyor.com/api/projects/nefarius/vigemclient/artifacts/bin/release/x64/ViGEmClient.dll?job=Platform%3A%20x64)
    - [x86](https://ci.appveyor.com/api/projects/nefarius/vigemclient/artifacts/bin/release/x86/ViGEmClient.dll?job=Platform%3A%20x86)

Place ViGEmClient.dll in the same folder as your application. The library should automatically detect dll on startup.

## References

- [ViGEm/ViGEmClient](https://github.com/ViGEm/ViGEmClient)
- [ViGEm/ViGEm.NET](https://github.com/ViGEm/ViGEm.NET)
- [yannbouteiller/vgamepad](https://github.com/yannbouteiller/vgamepad)
