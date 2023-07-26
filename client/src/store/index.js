import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

export function createStore (context) {
    return new Vuex.Store({
        state: {
            meta: {
                Title: "ShareUs",
                Keywords: "软件开发,web,html,javascript,css,ajax,golang,nodejs,java,c#,docker,ELK,vue,angular,react,SSR,CSR,koa,express,macos,linux,gogs,rabbitMQ,vi,vim,neovim,emacs,visual studio code,websocket,postgresql,数字签名,JSON Web Token,安全,VC,Nginx,HTTP,TCP/IP,DNS,Android,iOS,Mobile",
                Description: "技术博客，问题记录并提供解决方案，国外文章翻译，自创文章",
                OgImage: "https://github.githubassets.com/images/modules/site/logos/google-logo.png",
            },
            count: 0,
            pageData: {},
            xhrResult: "",
        },
        mutations: {
            increaseCount (state) {
                state.count++
            },  
            decreaseCount (state) {
                state.count--
            },
            setXhrResult(state, result) {
                state.xhrResult = result
            }
        },
        actions: {
            xhrTest ({commit}) {
                return axios.get('/api/check').then(res => {
                    console.log("XMLHttpRequest ok.")
                    commit('setXhrResult', res.data)
                })
            }
        }
    })
}
