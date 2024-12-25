# Use a lightweight Linux base image with PgBouncer
FROM debian:bullseye-slim

# Install PgBouncer
RUN apt-get update && apt-get install -y --no-install-recommends \
    pgbouncer postgresql-client \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /etc/pgbouncer

# Copy configuration files into the container
COPY pgbouncer.ini /etc/pgbouncer/
COPY userlist.txt /etc/pgbouncer/

# Create a PgBouncer user
RUN useradd -m -d /home/pgbouncer -s /bin/bash pgbouncer

# Make the log file and the PID file writable by the PgBouncer user
RUN mkdir -p /var/log/pgbouncer/
RUN mkdir -p /var/run/pgbouncer/

RUN touch /var/log/pgbouncer/pgbouncer.log
RUN touch /var/run/pgbouncer/pgbouncer.pid

RUN chown -R pgbouncer:pgbouncer /var/log/pgbouncer/
RUN chown -R pgbouncer:pgbouncer /var/run/pgbouncer/

RUN ls -ld /var/log/pgbouncer && ls -l /var/log/pgbouncer/pgbouncer.log

# Switch to the PgBouncer user
USER pgbouncer

# Expose PgBouncer's default port
EXPOSE 6432

# Command to start PgBouncer
CMD ["pgbouncer", "/etc/pgbouncer/pgbouncer.ini"]
