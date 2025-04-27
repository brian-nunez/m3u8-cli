# Command Line Tool for downloading `.m3u8` playlist videos

🎬 A simple, clean CLI tool to download `.m3u8` playlist videos using `ffmpeg`.

Built in Go. No dependencies. Just pure download power.

---

## 📦 Installation

Install directly using `go install`:

```sh
go install github.com/brian-nunez/m3u8-cli@latest
```

Make sure your ~/go/bin is in your $PATH to use m3u8-cli anywhere.

You must also have ffmpeg installed on your machine.

---

## 🚀 Usage

Basic download example:

```sh
m3u8-cli --url "https://example.com/playlist.m3u8" --output "my-video.mp4"
```

This will:

* Download the video from the .m3u8 playlist
* Save it as my-video.mp4 in the current directory

If you feel more comfortable with the terminal, you can use the help command to see all available options:

```sh
m3u8-cli --help
```

---

🛠️ Available Flags

| Flag           | Type   | Default      | Description                                 |
|:-------------- |:------ |:------------ |:------------------------------------------- |
| `--url`        | string | _(required)_ | The URL of the `.m3u8` playlist to download |
| `--output`     | string | `output.mp4` | Name of the output file                     |
| `--output-dir` | string | `.`          | Directory to save the output file           |
| `--timeout`    | int    | `30`         | Timeout (seconds) for stalled downloads     |
| `--quiet`      | bool   | `false`      | Suppress `ffmpeg` logs                      |
| `--version`    | bool   | `false`      | Print version info and exit                 |
| `--help`       | bool   |              | Show usage and available options            |

---

📝 Examples

Download a playlist quietly:

```sh
m3u8-cli --url "https://example.com/playlist.m3u8" --output "movie.mp4" --quiet
```

Save to a specific directory:

```sh
m3u8-cli --url "https://example.com/playlist.m3u8" --output "movie.mp4" --output-dir "/mnt/videos"
```

Check version:

```sh
m3u8-cli --version
```

---

⚡ Requirements

* Go 1.21+ to build
* ffmpeg installed and available in your system’s PATH

---

✨ License

See LICENSE file for details.

---

## Authors

- [brian-nunez](https://www.github.com/brian-nunez) - Maintainer

