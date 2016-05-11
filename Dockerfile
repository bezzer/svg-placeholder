FROM busybox
MAINTAINER Chris Berry <bezzer@gmail.com>
ADD www/ /www/
ADD svg-placeholder /app/svg-placeholder
ENV PORT 5000
EXPOSE 5000
ENTRYPOINT ["/app/svg-placeholder"]
