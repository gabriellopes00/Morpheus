import { LocalizationProvider, StaticDatePicker } from '@mui/lab'
import AdapterDateFns from '@mui/lab/AdapterDateFns'
import {
  Button,
  ButtonGroup,
  Checkbox,
  Divider,
  Grid,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Slider,
  TextField,
  ToggleButton,
  ToggleButtonGroup,
} from '@mui/material'
import { useState } from 'react'

export function SideBar() {
  const [value, setValue] = useState(null)

  function valuetext(value: number) {
    return `${value}°C`
  }

  const states = [
    { name: 'Acre (AC)', abbr: 'AC' },
    { name: 'Alagoas (AL)', abbr: 'AL' },
    { name: 'Amapá (AP)', abbr: 'AP' },
    { name: 'Amazonas (AM)', abbr: 'AM' },
    { name: 'Bahia (BA)', abbr: 'BA' },
    { name: 'Ceará (CE)', abbr: 'CE' },
    { name: 'Distrito Federal (DF)', abbr: 'DF' },
    { name: 'Espírito Santo (ES)', abbr: 'ES' },
    { name: 'Goiás (GO)', abbr: 'GO' },
    { name: 'Maranhão (MA)', abbr: 'MA' },
    { name: 'Mato Grosso (MT)', abbr: 'MT' },
    { name: 'Mato Grosso do Sul (MS)', abbr: 'MS' },
    { name: 'Minas Gerais (MG)', abbr: 'MG' },
    { name: 'Pará (PA)', abbr: 'PA' },
    { name: 'Paraíba (PB)', abbr: 'PB' },
    { name: 'Paraná (PR)', abbr: 'PR' },
    { name: 'Pernambuco (PE)', abbr: 'PE' },
    { name: 'Piauí (PI)', abbr: 'PI' },
    { name: 'Rio de Janeiro (RJ)', abbr: 'RJ' },
    { name: 'Rio Grande do Norte (RN)', abbr: 'RN' },
    { name: 'Rio Grande do Sul (RS)', abbr: 'RS' },
    { name: 'Rondônia (RO)', abbr: 'RO' },
    { name: 'Roraima (RR)', abbr: 'RR' },
    { name: 'Santa Catarina (SC)', abbr: 'SC' },
    { name: 'São Paulo (SP)', abbr: 'SP' },
    { name: 'Sergipe (SE)', abbr: 'SE' },
    { name: 'Tocantins (TO)', abbr: 'TO' },
  ]

  const [alignment, setAlignment] = useState('web')

  const handleChange = (
    event: React.MouseEvent<HTMLElement>,
    newAlignment: string
  ) => {
    setAlignment(newAlignment)
  }

  return (
    <Grid container gap={2}>
      <LocalizationProvider dateAdapter={AdapterDateFns}>
        <StaticDatePicker
          displayStaticWrapperAs="desktop"
          value={value}
          onChange={(newValue) => {
            setValue(newValue)
          }}
          renderInput={(params) => <TextField {...params} />}
        />
      </LocalizationProvider>

      <List
        sx={{
          width: '100%',
          maxWidth: 460,
          maxHeight: 460,
          bgcolor: 'background.paper',
          overflowY: 'scroll',
        }}
      >
        {states.map((state) => {
          const labelId = `checkbox-list-label-${state.name}`

          return (
            <ListItem key={state.abbr} disablePadding>
              <ListItemButton
                role={undefined}
                // onClick={handleToggle(state)}
                dense
              >
                <ListItemIcon>
                  <Checkbox
                    edge="start"
                    // checked={checked.indexOf(state) !== -1}
                    tabIndex={-1}
                    disableRipple
                    inputProps={{ 'aria-labelledby': labelId }}
                  />
                </ListItemIcon>
                <ListItemText id={labelId} primary={state.name} />
              </ListItemButton>
            </ListItem>
          )
        })}
      </List>

      <ToggleButtonGroup
        // color="primary"
        value={alignment}
        exclusive
        onChange={handleChange}
      >
        <ToggleButton value="0" color="success">
          L
        </ToggleButton>
        <ToggleButton value="10" color="primary">
          10
        </ToggleButton>
        <ToggleButton value="12" color="warning">
          12
        </ToggleButton>
        <ToggleButton value="14" color="error">
          14
        </ToggleButton>
        <ToggleButton value="16" color="info">
          16
        </ToggleButton>
        <ToggleButton value="18">18</ToggleButton>
      </ToggleButtonGroup>

      <Slider
        aria-label="maximum-capacity"
        defaultValue={30}
        valueLabelDisplay="auto"
        min={10}
        max={500}
      />
    </Grid>
  )
}
