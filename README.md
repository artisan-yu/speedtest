# speedtest

测速用例 - 生成指定大小的随机文件，用于磁盘/网络速度测试。

## 构建

```bash
go build -o gen .
```

## 用法

```bash
./gen <大小MB>
```

- 参数为正整数，单位为 MB
- 生成文件名为 `<大小>mb.txt`
- 文件内容为随机乱序的英文字母、数字、标点符号

## 示例

```bash
./gen 100    # 生成 100mb.txt (100 MB)
./gen 1      # 生成 1mb.txt (1 MB)
./gen 1024   # 生成 1024mb.txt (1 GB)
```

## 生成内容字符集

- 小写字母: `a-z`
- 大写字母: `A-Z`
- 数字: `0-9`
- 标点符号: `!@#$%^&*()-_=+[]{}|;:',.<>?/`~"\`

## 验证文件大小

```bash
# macOS
ls -la 100mb.txt

# Linux
stat 100mb.txt
```

## 发布gitPage测试socks5下行速率
```bash
curl --socks5 127.0.0.1:7897 \
  --max-time 8 -o /dev/null -s \
  -w "%{size_download} %{time_namelookup} %{time_connect} %{time_starttransfer} %{time_total} %{speed_download}\n" \
  https://speedtest.holi.vip/50mb.txt | \
awk '{
    mbs=$6/1024/1024;
    mbps=$6*8/1000000;
    if(mbps>=500) grade="★★★★★ 夯爆";
    else if(mbps>=200) grade="★★★★☆ 好用";
    else if(mbps>=50) grade="★★★☆☆ 够用";
    else if(mbps>=10) grade="★★☆☆☆ 能用";
    else if(mbps>=1) grade="★☆☆☆☆ 拉垮";
    else grade="☆☆☆☆☆ 垃圾";
    
    printf "\n📦 文件大小: %.2f MB\n", $1/1024/1024;
    printf "🌐 DNS解析: %.3fs\n", $2;
    printf "🤝 TCP连接: %.3fs\n", $3;
    printf "📥 首包时间: %.3fs\n", $4;
    printf "⏱️ 总耗时: %.3fs\n", $5;
    printf "🚀 平均速度: %.2f MB/s (%.0f Mbps)\n", mbs, mbps;
    printf "📊 线路评级: %s\n", grade;
}'
```
