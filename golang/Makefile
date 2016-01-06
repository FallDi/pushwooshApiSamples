CC=go
RESULT=pwCreateMessage
CFLAGS=build -gcflags "-N -l" -o
MAIN_FILE=createMessage.go

main: ${MAIN_FILE}
	${CC} ${CFLAGS} ${RESULT} ${MAIN_FILE}

clean:
	rm -f ${RESULT}