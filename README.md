# reddscare


```
reddscare [flags] <config-file> <subreddit>
```

reddscare saves reddit content to disk as html. It works by running several tools in sequence, so it is really just a glorified script.

## Installation

Download the go compiler, then run the below command:

```
go install github.com/ccollins476ad/reddscare@latest
```

Then install reddscare's three dependencies:
* [Bulk Downloader for Reddit](https://github.com/Serene-Arc/bulk-downloader-for-reddit)
* [bdfrscrape](https://github.com/ccollins476ad/bdfrscrape)
* [bdfrtohtml](https://github.com/BlipRanger/bdfr-html)

Finally, set up a reddscare config file. I recommend copying `resources/example.yaml` to a different location, then adjusting the path settings.

## Example

The below example downloads the AskHistorians subreddit to disk as html.

```
reddscare -v ~/tmp/reddscare.yaml AskHistorians
```

`~/tmp/reddscare.yaml` is just a copy of `resources/example.yaml`, which happens to work on my machine. After the command completes, open `/home/ccollins/tmp/html-test/AskHistorians/index.html` in your browser to view the downloaded content.
