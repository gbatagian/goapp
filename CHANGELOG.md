# [Unreleased] 2024/03/29

Initial version.

# [Unreleased] 2024/12/12

## Added

- Rule for image creation: `make image`
- Rule for container run: `make run`
- Rule for clean image build: `make clean-image`

## Changed

- Update `ENTRYPOINT` to `/goapp/server` to automatically start the server when the container runs
- Modify server binding to `0.0.0.0` to ensure the service is accessible from localhost while running in a container

## Fixed

- Session stats sent counter