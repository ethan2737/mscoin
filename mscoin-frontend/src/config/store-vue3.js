// Vuex 4.x for Vue 3 runtime store contract
import { runtimeContract } from './runtime-vue3'

const getStoredJson = (key) => {
    const value = localStorage.getItem(key);

    if (!value) {
        return null;
    }

    try {
        return JSON.parse(value);
    } catch (error) {
        localStorage.removeItem(key);
        return null;
    }
}

export default {
    state: () => ({
        host: runtimeContract.host,
        wshost: runtimeContract.wshost,
        api: runtimeContract.api,
        tokenKey: runtimeContract.tokenKey,
        memberKey: runtimeContract.memberKey,
        HeaderActiveName: '0',
        member: null,
        activeNav: '',
        lang: '',
        exchangeSkin: 'night',
        loginTimes: null
    }),
    mutations: {
        navigate(state, nav) {
            state.activeNav = nav;
        },
        removeMember(state) {
            state.member = null;
            localStorage.removeItem(state.memberKey);
        },
        setMember(state, member) {
            state.member = member;
            if (member == null) {
                localStorage.removeItem(state.memberKey);
                return;
            }
            localStorage.setItem(state.memberKey, JSON.stringify(member));
        },
        recoveryMember(state) {
            state.member = getStoredJson(state.memberKey);
        },
        setlang(state, lang) {
            state.lang = lang;
            localStorage.setItem('LANGUAGE', JSON.stringify(lang));
        },
        initLang(state) {
            if (localStorage.getItem('LANGUAGE') == null) {
                state.lang = '简体中文';
            } else {
                state.lang = JSON.parse(localStorage.getItem('LANGUAGE'));
            }
        },
        initLoginTimes(state) {
            if (localStorage.getItem('LOGINTIMES') == null) {
                state.loginTimes = 0;
            } else {
                state.loginTimes = JSON.parse(localStorage.getItem('LOGINTIMES'));
            }
        },
        setLoginTimes(state, times) {
            state.loginTimes = times;
            localStorage.setItem('LOGINTIMES', JSON.stringify(times));
        },
        setSkin(state, skin) {
            state.exchangeSkin = skin;
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
};
