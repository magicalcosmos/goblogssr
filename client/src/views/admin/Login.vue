<template>
  <div class="login">
    <div class="login-frame">
      <div class="left">
        <i class="icon icon-logo"></i>
      </div>
      <ul class="right">

        <li class="title-1"> {{ $t('login.already_member')}} </li>
        <li class="title-2">{{ $t('login.login_here')}}</li>
        <li class="line-top"></li>
        <li class="user">
          <input type="text" :placeholder="$t('common.email')"/>
        </li>
        <li class="pwd">
          <input type="password" :placeholder="$t('common.password')"/>
        </li>
        <li class="opt">
          <span class="me">
            <input type="checkbox"/>
            <span class="rember">{{ $t('login.rember_me')}}</span>
          </span>
          <span class="buttons">
            <button class="reset">
              <span class="rember">{{ $t('button.reset')}}</span>
            </button>
            <button>
              <span class="rember">{{ $t('button.submit')}}</span>
            </button>
          </span>
        </li>
        <li class="line-bottom"></li>
      </ul>
    </div>
  </div>
</template>
<script>
  export default {
    data() {
      return {};
    },
    methods: {
      
    }
  }
</script>
<style scoped>
.login {
  background: url('@static/img/bg_login.png') repeat;
  background-size: contain;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  width: 100%;
}

.login-frame .left {
  float: left;
  width: 58px;
}

.login-frame .left .icon-logo {
  background: url('@static/img/icon_logo.png') no-repeat;
  height: 34px;
  width: 58px;
}

.login-frame .right {
  float: left;
  margin-left: 10px;
  width: 235px;

}

.login-frame .right input[type="text"],
.login-frame .right input[type="password"] {
  outline: none;
  border: 1px solid #e0e3e3;

  background: linear-gradient(to bottom, #f1f1f1 50%, #eff0F0 100%);
  box-shadow: 0px 2px 0px 0px #ced0d0;
  height: 28px;
  opacity: 1;
  padding: 0 7px;
  width: 100%;
}

.login-frame .right input[type="text"]::placeholder,
.login-frame .right input[type="password"]::placeholder {
  font-size: 12px;
  font-family: HelveticaLT-Condensed, HelveticaLT;
  font-weight: normal;
  color: #c3c5c5;
  line-height: 0px;
  text-shadow: 1px 0px 0px rgba(255, 255, 255, .5);
}


.login-frame .right input[type="checkbox"] {
	-moz-appearance:none;
	-webkit-appearance:none;
	-o-appearance:none;

  box-shadow: -2px -1px 1px 1px rgba(0, 0, 0, .2);
	content: none;	
  height: 16px;
  margin-left: 2px;
	outline: none;
  position: relative;
  top: 2px;
  width: 16px;
}

.login-frame .right input[type="checkbox"]:before {
  border: none;
  color: transparent !important;
  content: "\f00c";
  display: block;
  font-family: "FontAwesome";
  font-size: 15px;
  height: 16px;
  width: 16px;
}

.login-frame .right input[type="checkbox"]:checked:before {
  color: black !important;
}

.login-frame .right .title-1 {
  color: #8d8d8d;
  font-size: 18px;
  font-family: HelveticaLT-Condensed, HelveticaLT;
  font-weight: normal;
  text-shadow: 1px 0px 0px #fff;
  margin-top: -20px;
}

.login-frame .right .title-2 {
  color: #7B7B7B;
  font-family: HelveticaLT-Condensed, HelveticaLT;
  font-size: 28px;
  font-weight: normal;
  text-shadow: 1px 0px 0px #fff;
}

.login-frame .right .line-top,
.login-frame .right .line-bottom {
  border-top: 2px solid rgba(0, 0, 0, 0.18);
  border-top-left-radius: 2px;
  border-top-right-radius: 2px;
  box-shadow: 0px 1px 0px 0px rgba(255, 255, 255, 0.5);
  opacity: 1;
  height: 1px;
  margin-top: 12px;
  opacity: 1;
  width: 235px;
}

.login-frame .right .line-bottom {
  border-top: 2px solid rgba(0, 0, 0, 0.1);
}

.login-frame .right li {
  width: 235px;
}

.login-frame .right .user,
.login-frame .right .pwd {
  margin-top: 8px;
}

.login-frame .right .opt {
  margin-top: 15px;
  overflow: hidden;
}


.login-frame .right .opt button {
  background-image: linear-gradient(to bottom, #f1f1f1 50%, #dcdddf 100%);
  border: 1px solid #e0e3e3;
  box-shadow: 1px 0px 0px 0px #ced0d0;
  height: 28px;
  line-height: 28px;
  cursor: pointer;
  opacity: 1;
  text-align: center;
  width: 50px;
	filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#fa5a5a', endColorstr='#e85151', GradientType=0 );
}

.login-frame .right .opt button:hover {
  background-image: linear-gradient(to bottom, #f7f7f7 50%, #dcdddf 100%);
  border: 1px #c6c6c6 solid;
  box-shadow: 1px 1px 1px #eaeaea;
  color: #333;
}

.login-frame .right .opt button:active {
  box-shadow: inset 1px 1px 1px #DFDFDF; 
}

.login-frame .right .opt .rember {
  color: #7b7b7b;
  font-size: 12px;
  font-family: HelveticaLT-Condensed, HelveticaLT;
  font-weight: normal;
  text-shadow: 1px 0px 0px #fff;
}

.login-frame .right .opt .me {
  float: left;
}

.login-frame .right .opt .buttons {
  float: right;
}

.login-frame .right .opt .reset {
  margin-right: 5px;
}

</style>
