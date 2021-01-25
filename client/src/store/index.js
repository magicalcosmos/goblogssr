import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

export function createStore (context) {
    return new Vuex.Store({
        state: {
            meta: {
                Title: "ShareUs",
                Keywords: "Love is meant to be shared – https://shareus.cn",
                Description: "排序,正则表达式,程序员,算法,软件开发,ajax,algorithm,android,bash,book,c++,code review,coding,css,database,debug,ebook,erlang,game,go,google,html,ie,java,javascript,jquery,linus torvalds,linux,mac,mysql,oracle,os,perl,php,programmer,programming,language,python,ruby,sql,ubuntu,ui,unix,vim,web,windows",
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
