import Home from './pages/Home'
const Scan = () => import('./pages/Scan')
export default [
  {path:'/', component: Home},
  {path:'/scan', component: Scan}
]
