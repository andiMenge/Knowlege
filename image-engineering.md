# Transcoding

## h.264 VBV Buffer
The Video Buffer Verifier (VBV) is a model hypothetical decoder buffer that will not overflow or underflow when fed a conforming MPEG bit stream.

```vbv_buffer_size_value = the maximum buffer fullness```

```vbv_delay = the delay between storing a picture start code in the buffer and starting the decoding of that picture```

Bufsize specifies the size of the device's buffer

Very basically, VBV keeps track of the buffer filling level

bufsize / maxrate = number of seconds of wall clock time the decoder has to buffer before it starts playing.

max_buffering_delay (second) = buffer_size (bytes) * 8 / bit_rate (bits/second)
https://codesequoia.wordpress.com/2010/04/19/buffering-delay-and-mpeg-2-transport-stream/

## FFMPEG -bufsize
-bufsize which tells the encoder how often to calculate the average bit rate and check to see if it conforms to the average bit rate specified on the command line (-b:v )

## FFMPEG Interlaced h.264@TS CBR, fixed GOP

```ffmpeg -i in -c:v libx264 -preset slow -c:a copy -b:v 2000k -minrate 2000k -maxrate 2000k -r 25 -g 25 -bufsize 2000k -nal-hrd cbr -x264-params scenecut=-1:interlaced=1 -t 00:05:00 out.ts```

## GoPro Timelapse to Video
`cd goPro image folder`

### With upside down correction
`find -s . -type f -name *.JPG  -exec cat {} ';' | ffmpeg -f image2pipe -r 12 -i - -c:v h264 -preset ultrafast -pix_fmt yuv420p -s hd720 -vf "vflip,hflip" -r 24 foo.mp4`
