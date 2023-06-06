# VRM Version Viewer

VRMとUniVRMのバージョン見るやーつ

## 使い方

```
./vrm_version_viewer-x_x_x-amd64.exe avater.vrm
```


## test
```bash
docker compose run --rm builder bash

go run vrm_version_viewer.go VRM_FILE
```

# 参考

- [VRM 0.x specification](https://github.com/vrm-c/vrm-specification/tree/master/specification/0.0)
- [VRM 1.0 specification](https://github.com/vrm-c/vrm-specification/tree/master/specification/VRMC_vrm-1.0)
- [GLTF 2.0 specification](https://github.com/KhronosGroup/glTF/tree/main/specification/2.0)

# License
MIT
