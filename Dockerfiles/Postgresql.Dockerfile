FROM postgres:16 as basic

RUN apt-get update -y --no-install-recommends

RUN apt-get install -y --no-install-recommends \
	gcc \
	make

RUN apt-get install -y --no-install-recommends pgxnclient && \
	pgxn install ddlx

FROM basic as runtime

EXPOSE 5432

CMD ["postgres"]
