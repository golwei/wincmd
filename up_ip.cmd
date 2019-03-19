(for /f "tokens=1 skip=2"  %%i in ('netsh interface ipv4 show interface') do (netsh interface ip set address name=%%i  source=dhcp))
