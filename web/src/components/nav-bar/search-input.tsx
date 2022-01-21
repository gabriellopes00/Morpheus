import { Search as SearchIcon } from '@mui/icons-material'
import {
  AutocompleteRenderInputParams,
  Box,
  InputBase,
  useTheme,
} from '@mui/material'
import { alpha, styled } from '@mui/material/styles'

export interface SearchInputProps {
  params: AutocompleteRenderInputParams
}

export function SearchInput({ params }: SearchInputProps) {
  const IconWrapper = styled('div')(({ theme }) => ({
    padding: theme.spacing(0, 2),
    height: '100%',
    position: 'absolute',
    pointerEvents: 'none',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  }))

  const TextInput = styled(InputBase)(({ theme }) => ({
    color: 'inherit',
    '& .MuiInputBase-input': {
      padding: theme.spacing(1, 1, 1, 0),
      paddingLeft: `calc(1em + ${theme.spacing(4)})`,
      transition: theme.transitions.create('width'),
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        width: '24ch',
        '&:focus': { width: '34ch' },
      },
      [theme.breakpoints.up('lg')]: {
        width: '44ch',
        '&:focus': { width: '64ch' },
      },
    },
  }))

  const theme = useTheme()

  return (
    <Box
      ref={params.InputProps.ref}
      sx={{
        position: 'relative',
        borderRadius: '4px',
        backgroundColor: alpha(theme.palette.common.white, 0.15),
        '&:hover': { backgroundColor: alpha(theme.palette.common.white, 0.25) },
        marginLeft: 0,
        width: '100%',
        [theme.breakpoints.up('sm')]: {
          marginLeft: theme.spacing(1),
          width: 'auto',
        },
      }}
    >
      <IconWrapper>
        <SearchIcon />
      </IconWrapper>
      <TextInput
        placeholder="Search…"
        inputProps={{ 'aria-label': 'search', ...params.inputProps }}
      />
    </Box>
  )
}
