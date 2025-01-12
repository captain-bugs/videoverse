# videoverse
 
## Overview

Videoverse is a video management and processing application

To get started you can either run it locally or use Docker for a containerized deployment.
Follow the instructions in the "Getting Started" section to set up the application.

Link to Demo [demo](https://drive.google.com/file/d/1Q10n3bp_JaUMlz6iPummt72G96OS7MWF/view?usp=sharing)

## Features

- **User Management**: Create and manage user accounts.
- **Video Upload**: Upload videos with metadata.
- **Video Trimming**: Trim videos to create new clips.
- **Video Merging**: Merge multiple videos into one.
- **Video Sharing**: Generate shareable links for videos.
- **Authentication**: Secure endpoints with JWT-based authentication.

## Project Structure

- **cmd**: Contains the main application entry point and API handlers.
- **pkg**: Contains various packages for configuration, authentication, models, and utilities.
- **repository**: Contains the database repository implementations.
- **routes**: Contains the route definitions for the application.
- **db**: Contains database schema and migration files.
- **av**: Contains video processing logic using ffmpeg.
- **middleware**: Contains middleware for authentication and other purposes.
- **storage**: Contains file storage implementations.

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Docker (optional, for containerized deployment)
- FFmpeg (for video processing)


###### (Optional) If you prefer using Docker, simply run the following command to start the application:
```sh
make docker-run
```

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/captain-bugs/videoverse.git
    cd videoverse
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

3. Run database migrations:
    ```sh
    make migrate-up
    ```

4. Build the project:
    ```sh
    make build
    ```

5. Run the application:
    ```sh
    ./bin/videoverse
    ```


## API Endpoints

Postman collection can be found in `docs` folder

- **User Endpoints**
  - `POST /api/v1/user/`: Create a new user.
  - `GET /api/v1/user/`: Get user details (requires authentication).

- **Video Endpoints**
  - `POST /api/v1/video/`: Upload a new video (requires authentication).
  - `GET /api/v1/video/list/`: List user videos (requires authentication).
  - `GET /api/v1/video/:id/`: Get video details (requires authentication).
  - `POST /api/v1/video/trim/`: Trim a video (requires authentication).
  - `POST /api/v1/video/merge/`: Merge videos (requires authentication).

- **Share Endpoints**
  - `GET /api/v1/share/video/:id/`: Generate a shareable link for a video (requires authentication).
  - `GET /api/v1/share/view/`: View a shared video using the shareable link.

## Configuration

Configuration is managed through environment variables. The following variables can be set:

- `ENV`: Application environment (`dev` or `production`).
- `APP_PORT`: Port on which the application runs.
- `JWT_SECRET`: Secret key for JWT authentication.
- `DATABASE_PATH`: Path to the SQLite database file.
- `FILE_UPLOAD_PATH`: Path for storing uploaded video files.
- `CDN_ENDPOINT`: Endpoint for accessing shared video links.


## Acknowledgements

- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP web framework.
- [FFmpeg](https://ffmpeg.org/) - Video processing library.
- [sqlc](https://github.com/kyleconroy/sqlc) - Generate type-safe Go from SQL.

