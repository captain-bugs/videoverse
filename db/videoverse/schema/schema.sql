CREATE TABLE IF NOT EXISTS users
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    username      TEXT NOT NULL UNIQUE,
    email         TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS videos
(
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id         INTEGER NOT NULL,                                                   -- References the user who owns this video
    source_video_id INTEGER,                                                            -- References the original video if this is a trimmed or merged video
    type            TEXT    NOT NULL CHECK (type IN ('ORIGINAL', 'TRIMMED', 'MERGED')), -- Type of video
    file_path       TEXT    NOT NULL,                                                   -- Path to the video file
    file_name       TEXT,                                                               -- Original file name for 'original' videos
    size_in_bytes   INTEGER NOT NULL,                                                   -- Size of the video in bytes
    duration        INTEGER NOT NULL,                                                   -- Duration of the video in seconds
    start_time      INTEGER,                                                            -- Start time for trimmed videos
    end_time        INTEGER,                                                            -- End time for trimmed videos
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME,
    FOREIGN KEY (source_video_id) REFERENCES videos (id) ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE IF NOT EXISTS shared_links
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    video_id   INTEGER  NOT NULL,        -- References videos(id)
    user_id    INTEGER  NOT NULL,        -- References users(id) who shared the link
    link       TEXT     NOT NULL UNIQUE, -- Unique link for sharing
    expires_at DATETIME NOT NULL,        -- Expiry time for the link
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (video_id) REFERENCES videos (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id)
    );
