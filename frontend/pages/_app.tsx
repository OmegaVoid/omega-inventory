import '../styles/globals.css'
import type { AppProps } from 'next/app'
import {library} from '@fortawesome/fontawesome-svg-core'
import {
  faArrowAltCircleLeft,
  faEdit,
  faFile,
  faFolder,
  faFolderMinus,
  faFolderOpen,
  faFolderPlus,
  faHome,
  faList,
  faMinusSquare,
  faPlusSquare,
  faSignOutAlt,
  faSync
} from '@fortawesome/free-solid-svg-icons'
import React from "react";

library.add(faFolder, faFolderOpen, faFile, faFolderPlus, faFolderMinus,
    faPlusSquare, faMinusSquare, faEdit, faSync, faHome, faList, faSignOutAlt, faArrowAltCircleLeft)


function MyApp({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />
}
export default MyApp
