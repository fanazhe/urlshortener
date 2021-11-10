import { MutationTree } from 'vuex'
import ShortenModel from './models/ShortenModel'
import { State } from './state'

const LOCALSTORAGE_URL_COUNT: number = +import.meta.env.VITE_LOCALSTORAGE_URL_COUNT;

export enum Mutation {
  NEW_SHORTEN = 'NEW_SHORTEN',
  TOGGLE_ABOUT = 'TOGGLE_ABOUT',
  TOGGLE_STATS = 'TOGGLE_STATS',
  TOGGLE_LOGIN = 'TOGGLE_LOGIN',
}

export type Mutations<S = State> = {
  [Mutation.NEW_SHORTEN](state: S, payload: ShortenModel): void,
  [Mutation.TOGGLE_ABOUT](state: S, payload: undefined): boolean,
  [Mutation.TOGGLE_STATS](state: S, payload: undefined): boolean,
  [Mutation.TOGGLE_LOGIN](state: S, payload: undefined): boolean,
}

export const mutations: MutationTree<State> & Mutations = {
  [Mutation.NEW_SHORTEN](state: State, payload: ShortenModel) {
    state.shortens.push(payload);
    while (state.shortens.length > LOCALSTORAGE_URL_COUNT) {
      state.shortens.shift();
    }
  },
  [Mutation.TOGGLE_ABOUT](state: State, payload: undefined) {
    state.showAbout = !state.showAbout;
    return state.showAbout;
  },
  [Mutation.TOGGLE_STATS](state: State, payload: undefined) {
    state.showStats = !state.showStats;
    return state.showStats;
  },
  [Mutation.TOGGLE_LOGIN](state: State, payload: undefined) {
    state.showLogin = !state.showLogin;
    return state.showLogin;
  },
}
