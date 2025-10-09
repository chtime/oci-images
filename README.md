# OCI Images

A repository for building and maintaining custom OCI (Open Container Initiative) images using GitHub Actions.

## Structure

This repository uses a dynamic folder-based approach where each folder represents a separate OCI image to be built. The GitHub Actions workflow automatically discovers all folders containing a `Containerfile` and builds them as separate images.

### Folder Structure

Each image folder should contain:
- `Containerfile` - The container build instructions
- Any additional files, scripts, or resources needed by the image

### Current Images

- **nginx-custom** - A customized Nginx server with security headers and health check endpoint

## Usage

### Adding a New Image

1. Create a new folder in the root of the repository
2. Add a `Containerfile` to define the image build process
3. Include any additional files needed by your container
4. Commit and push - the GitHub Actions workflow will automatically build your image

### Built Images

Images are automatically built and pushed to GitHub Container Registry (ghcr.io) with the following naming convention:
```
ghcr.io/<owner>/<repo>/<image-folder-name>:<tag>
```

For example:
- `ghcr.io/chtime/oci-images/nginx-custom:latest`
- `ghcr.io/chtime/oci-images/nginx-custom:main`
- `ghcr.io/chtime/oci-images/nginx-custom:sha-abc123`

### Running Images

```bash
# Run the nginx-custom image
docker pull ghcr.io/chtime/oci-images/nginx-custom:latest
docker run -p 8080:80 ghcr.io/chtime/oci-images/nginx-custom:latest
```

## GitHub Actions Workflow

The `.github/workflows/build-images.yml` workflow:

1. **Discovery Phase**: Scans the repository for folders containing `Containerfile`
2. **Build Phase**: Uses a matrix strategy to build each discovered image in parallel
3. **Push Phase**: Pushes built images to GitHub Container Registry with appropriate tags

The workflow triggers on:
- Push to main/master branch
- Pull requests to main/master branch
- Manual workflow dispatch

## Contributing

To add a new image:

1. Create a folder with a descriptive name
2. Add a `Containerfile` with your image definition
3. Include any supporting files your image needs
4. Test your Containerfile locally if possible
5. Submit a pull request

The automated workflow will validate and build your image upon merge.