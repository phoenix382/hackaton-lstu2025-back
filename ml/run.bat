@echo off
start "server" cmd /k python server.py
timeout /t 1

@REM go\bin\go run client.go
@REM rem start "client1" cmd /k python client1.py


set err=%errorlevel%
IF NOT %err%==0 (pause)
pause
exit