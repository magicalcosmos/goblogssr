<template>
  <article>
    <p>
      <a href="#/blog">← 返回</a>
    </p>
    <hr>
    <header>
      <h1>[System] Mac系统的环境变量加载顺序</h1>
      <div class="author">
        <span>所属目录：Frontend</span>
        <span>创建者：Brody</span>
        <span>创建时间：2012-01-01</span>
      </div>
    </header>
    <div class="post">
      Mac系统的环境变量，加载顺序为：
/etc/profile /etc/paths ~/.bash_profile ~/.bash_login ~/.profile ~/.bashrc

1）/etc/paths （全局建议修改这个文件 ）
编辑 paths，将环境变量添加到 paths文件中 ，一行一个路径
Hint：输入环境变量时，不用一个一个地输入，只要拖动文件夹到 Terminal 里就可以了。

2）/etc/profile （建议不修改这个文件 ）
全局（公有）配置，不管是哪个用户，登录时都会读取该文件。

3）/etc/bashrc （一般在这个文件中添加系统级环境变量）
全局（公有）配置，bash shell执行时，不管是何种方式，都会读取此文件。

4）
1.创建一个文件：
sudo touch /etc/paths.d/mysql
2.用 vim 打开这个文件（如果是以 open -t 的方式打开，则不允许编辑）：
sudo vim /etc/paths.d/mysql
3.编辑该文件，键入路径并保存（关闭该 Terminal 窗口并重新打开一个，就能使用 mysql 命令了）
/usr/local/mysql/bin
据说，这样可以自己生成新的文件，不用把变量全都放到 paths 一个文件里，方便管理。

5) ~/.bash_profile （任意一个文件中添加用户级环境变量）

注：Linux 里面是 .bashrc 而 Mac 是 .bash_profile）
    </div>
  </article>
</template>
<script>
export default {
  
}
</script>
<style scoped>
  h1 {
    font-size: 36px;
  }
  p {
    margin-top: 20px;
  }

  p a {
    text-decoration: none;
  }

  p a:hover {
    text-decoration: underline;
  }

  hr {
    border: 0;
    border-top: 1px solid #eee;
    margin-top: 20px;
    margin-bottom: 20px;
  }

  .author {
    font-size: 12px;
    margin-top: 5px;
  }

  .author span {
    color: rgba(0, 0, 0, 0.5);
    margin-right: 10px;
  }

  .post {
    white-space: pre-wrap;
  }
</style>
