# 使用scratch开启
FROM scratch
# 拷贝编译程序
COPY main main
# 打开8080端口
EXPOSE 8080
# 运行!
CMD ["./main"]