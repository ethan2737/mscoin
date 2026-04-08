// Vuex 3.x 在 Vue 3 中使用 - 不使用 Vue.use()
// Vuex 3.x 支持 Vue 3，只需要直接导出 store 实例即可
import Vuex from 'vuex'

export default new Vuex.Store({
    state: {
        member: null,
        activeNav: '',
        lang: '',
        exchangeSkin:'night',
        loginTimes: null
    },
    mutations: {
        navigate(state, nav) {
            state.activeNav = nav;
        },
        removeMember(state) {
            state.member = null;
            localStorage.removeItem("MEMBER");
        },
        setMember(state, member) {
            state.member = member;
            localStorage.setItem('MEMBER', JSON.stringify(member));
        },
        recoveryMember(state) {
            state.member = JSON.parse(localStorage.getItem('MEMBER'));
        },
        setlang(state, lang) {
            state.lang = lang;
            localStorage.setItem('LANGUAGE', JSON.stringify(lang));
        },
        initLang(state) {
            if (localStorage.getItem('LANGUAGE') == null) {
                state.lang = "简体中文";
            } else {
                state.lang = JSON.parse(localStorage.getItem('LANGUAGE'));
            }
        },
        initLoginTimes(state){
            if(localStorage.getItem("LOGINTIMES") == null){
                state.loginTimes = 0;
            }else{
                state.loginTimes = JSON.parse(localStorage.getItem('LOGINTIMES'));
            }
        },
        setLoginTimes(state, times){
            state.loginTimes = times;
            localStorage.setItem('LOGINTIMES', JSON.stringify(times));
        },
        setSkin(state,skin){
            state.exchangeSkin=skin;
        }
    },
    getters: {
        member(state) {
            return state.member;
        },
        isLogin(state) {
            return state.member != null;
        },
        lang(state) {
            return state.lang;
        },
        loginTimes(state) {
            return state.loginTimes;
        }
    }
});
