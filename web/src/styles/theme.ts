import { createTheme } from '@mui/material/styles'
import { red } from '@mui/material/colors'

// Create a theme instance.
const theme = createTheme({
  palette: {
    primary: {
      main: '#14213d',
    },
    secondary: {
      main: '#fca311',
    },
    error: {
      main: red.A400,
    },
    mode: 'light',
    background: {
      default: '#e5e5e5',
    },
  },
})

export default theme
