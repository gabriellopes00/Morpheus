import {
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  Checkbox,
  ListItemText,
} from '@mui/material'
import { useState } from 'react'

export function StatesList() {
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

  const [selectedStates, setSelectedStates] = useState<string[]>([])

  return (
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
              onClick={() => setSelectedStates([...selectedStates, state.abbr])}
              dense
            >
              <ListItemIcon>
                <Checkbox
                  edge="start"
                  checked={selectedStates.indexOf(state.abbr) !== -1}
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
  )
}
