(for /f "tokens=1 skip=2"  %%i in ('netsh interface ipv4 show interface') do (netsh interface ip set address name=%%i  source=dhcp))
## 定时任务
schtasks /create /sc hourly /mo 5 /sd 2019/04/20 /tn "My App" /tr .\aaa.bat
