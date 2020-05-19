# テスト画像について

## 連番画像 (test-img%02d.png)

```
ffmpeg -f lavfi -t 10 -i testsrc=duration=10:rate=1 test-img%02d.png
```

## 連番アニメ (test-img.gif)

```
ffmpeg -f lavfi -t 10 -r 2 -i testsrc=duration=10:rate=1 test-img.gif
```

