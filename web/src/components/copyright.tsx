import { Typography } from '@mui/material'
import Link from '@mui/material/Link'

export default function Copyright(props: any) {
  return (
    <Typography
      variant="body2"
      color="text.secondary"
      align="center"
      {...props}
    >
      {'Copyright © '}
      <Link color="inherit" href="https://mrpheus.vercel.app">
        Morpheus
      </Link>{' '}
      {new Date().getFullYear() !== 2021 && '2021'} - {new Date().getFullYear()}
      {'.'}
    </Typography>
  )
}
