## Template for GitHub Apps built with Golang

### Running

```bash
docker run --rm --name go-github-app \
    -v $(pwd)/config:/config \
    -p 8080:8080 \
    ghcr.io/martinheinz/go-github-app/app
```