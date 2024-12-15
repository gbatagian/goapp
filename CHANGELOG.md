# [Unreleased] 2024/03/29

Initial version.

# [Unreleased] 2024/12/12

## Added

- Rule for image creation: `make image`
- Rule for container run: `make run`
- Rule for clean image build: `make clean-image`
- `/bin` on `.gitignore`
- `value` on web socket response message
- Test cases and benchmark for `RandString` function

## Changed

- Update `ENTRYPOINT` to `/goapp/server` to automatically start the server when the container runs
- Modify server binding to `0.0.0.0` to ensure the service is accessible from localhost while running in a container

## Fixed

- Session stats sent counter

# [Unreleased] 2024/12/13

## Added

- Command line client
- Semaphores to manage the maximum number of active WebSockets

# [Unreleased] 2024/12/14

## Added

- Gracefully handle errors on cli client
- Fine tune cli client log message

# [Unreleased] 2024/12/15

- Add allowed origins check on web socket connections
