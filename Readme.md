#### Casbin
使用casbin有两种权限校验的方法：
1. 不使用role_definition，在middleware中用token中的userid去DB里换到其role，然后用casbin去校验。
优点：不需要根据用户角色的更改而同步更改casbin中的数据。
缺点：每次请求都要去DB换取role。
*2. 使用role_definition，middleware只需要username，就可以自动判断权限。
优点：每次请求可以节省验证权限的性能开销。
缺点：每一次的用户角色的改变必须同步给casbin内置的结构表。