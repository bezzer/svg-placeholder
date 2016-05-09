FROM scratch
MAINTAINER Chris Berry <bezzer@gmail.com>
ADD svg-placeholder /app/svg-placeholder
ENV PORT 5000
EXPOSE 5000
ENTRYPOINT ["/app/svg-placeholder"]
