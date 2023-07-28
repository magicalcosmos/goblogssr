<template>
  <ul class="contact">
    <li>
      <div>
        <label for="username">{{ $t('contact.username') }}</label>
      </div>
      <div>
        <input type="text" name="name"/>
      </div>
    </li>
    <li>
    <li>
      <div>
        <label for="email">{{ $t('common.email') }}</label>
      </div>
      <div>
        <input type="text" name="name" />
      </div>
    </li>

    <li>
      <div>
        <label for="phone">{{ $t('contact.phone') }}</label>
      </div>
      <div>
        <input type="text" name="phone" :placeholder="$t('contact.placeholder.optional')"/>
      </div>
    </li>

    <li>
      <div>
        <label for="why">{{ $t('contact.why') }}</label>
      </div>
      <div>
        <select>
          <option value="">({{  $t('common.select_one') }})</option>
          <option value="1">{{  $t('contact.options.only_message') }}</option>
          <option value="2">{{  $t('contact.options.question') }}</option>
          <option value="3">{{  $t('contact.options.other') }}</option>
        </select>
      </div>
    </li>

    <li>
      <div>
        <label for="message">{{ $t('contact.message') }}</label>
      </div>
      <div>
        <textarea :placeholder="$t('contact.placeholder.leave_message')"></textarea>
      </div>
    </li>
    <li>
      <button class="btn">{{ $t('common.send')}}</button>
    </li>
  </ul>
</template>
<script>
  export default {
    
  };
</script>
<style scoped>
  ul {
    margin-top: 40px;
  }
  @media (min-width: 768px) {
    ul {
      width: 66.66667%;
    }
  }

  @media (min-width: 992px) {
    ul {
      width: 50%;
    }
  }

  li {
    margin-top: 15px;
  }

  li > div:last-child {
    margin-top: 5px;
  }

  input, select, textarea {
    display: block;
    width: 100%;
    height: 34px;
    padding: 6px 12px;
    font-size: 14px;
    line-height: 1.42857;
    color: #555555;
    background-color: #fff;
    background-image: none;
    border: 1px solid #ccc;
    border-radius: 4px;
    -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
    box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
    -webkit-transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
    -o-transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
    transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
  }

  textarea {
    height: auto;
  }

  button {
    background-color: #337ab7;
    border-color: #2e6da4;
    border-radius: 4px;
    color: #fff;
    padding: 6px 12px;
  }
</style>
