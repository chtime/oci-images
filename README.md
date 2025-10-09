# OCI Images

A repository for building and maintaining custom OCI (Open Container Initiative) images using GitHub Actions.

## Structure

This repository uses a dynamic folder-based approach where each folder represents a separate OCI image to be built. The GitHub Actions workflow automatically discovers all folders containing a `Containerfile` and builds them as separate images.

### Folder Structure

Each image folder should contain:
- `Containerfile` - The container build instructions
- Any additional files, scripts, or resources needed by the image

