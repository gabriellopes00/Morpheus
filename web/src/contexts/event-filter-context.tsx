import { createContext, ReactNode, useContext, useState } from 'react'

export interface EventFilterContextData {
  filter: EventFilter
  updateFilter: (filter: EventFilter) => void
}

export const EventFilterContext = createContext<EventFilterContextData>(
  {} as EventFilterContextData
)

export interface EventFilter {
  ageGroup: number
  states: string[]
  date: Date
  maximumCapacity: number
}

interface EventFilterProviderProps {
  children: ReactNode
}

export function EventFilterProvider({ children }: EventFilterProviderProps) {
  const [eventFilter, setEventFilter] = useState<EventFilter>({
    states: [],
    date: new Date(),
    ageGroup: 0,
    maximumCapacity: 200,
  })

  function updateFilter(filter: EventFilter): void {
    setEventFilter(filter)
  }

  return (
    <EventFilterContext.Provider value={{ filter: eventFilter, updateFilter }}>
      {children}
    </EventFilterContext.Provider>
  )
}

export function useFilter() {
  return useContext(EventFilterContext)
}
