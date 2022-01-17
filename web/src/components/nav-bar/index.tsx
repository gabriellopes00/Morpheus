import { AppBar, Box, Container, Toolbar } from '@mui/material'
import { AvatarMenu } from './avatar'
import { NavMenu } from './menu'
import { SearchBox } from './search-box'

export default function NavBar() {
  return (
    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <NavMenu />

          <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
            <SearchBox />
          </Box>

          <Box sx={{ flexGrow: 0 }}>
            <AvatarMenu />
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  )
}
