import ShortenModel from './models/ShortenModel'

export interface State {
  debug: boolean
  version: string
  isInitialized: boolean
  showAbout: boolean
  showStats: boolean
  showLogin: boolean
  shortens: ShortenModel[]
}

const versionString =
  import.meta.env.MODE === 'development'
    ? import.meta.env.VITE_APP_VERSION + '-dev'
    : import.meta.env.VITE_APP_VERSION

export const state: State = {
  debug: import.meta.env.MODE === 'development',
  version: versionString,
  isInitialized: false,
  showAbout: false,
  showStats: false,
  showLogin: false,
  shortens: Array<ShortenModel>(),
}
