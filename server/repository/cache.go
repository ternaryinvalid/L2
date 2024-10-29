package repository

import (
	"L2/server/models"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.Mutex
	data map[string][]*models.Event
}

func NewCache() *Cache {
	return &Cache{
		mu:   sync.Mutex{},
		data: make(map[string][]*models.Event),
	}
}

func (c *Cache) GetDay(date string) []*models.Event {
	c.mu.Lock()
	defer c.mu.Unlock()

	events, ok := c.data[date]

	if !ok {
		return []*models.Event{}
	}

	return events
}

func (c *Cache) Create(event *models.Event) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[event.Date] = append(c.data[event.Date], event)

}

func (c *Cache) Delete(date, time string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := 0; i < len(c.data[date]); i++ {
		if c.data[date][i].Time == time {
			c.data[date] = append(c.data[date][:i], c.data[date][i+1:]...)
		}
	}

}

func (c *Cache) GetWeek(date string) [][]*models.Event {
	var week = 168 * time.Hour
	parsedDate, _ := time.Parse("2006-01-02", date)

	timeSinceWeek := parsedDate.Add(week)
	var response [][]*models.Event
	c.mu.Lock()
	defer c.mu.Unlock()

	for iterDate, events := range c.data {
		parsedIter, _ := time.Parse("2006-01-02", iterDate)
		if parsedIter.Before(timeSinceWeek) && parsedIter.After(parsedDate) {
			response = append(response, events)
		}
	}
	return response
}

func (c *Cache) GetMonth(date string) [][]*models.Event {
	var month = 4 * 168 * time.Hour
	parsedDate, _ := time.Parse("2006-01-02", date)

	timeSinceMonth := parsedDate.Add(month)
	var response [][]*models.Event
	c.mu.Lock()
	defer c.mu.Unlock()

	for iterDate, events := range c.data {
		parsedIter, _ := time.Parse("2006-01-02", iterDate)
		if parsedIter.Before(timeSinceMonth) && parsedIter.After(parsedDate) {
			response = append(response, events)
		}
	}
	return response
}

func (c *Cache) Update(date, time string, newEvent *models.Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.data[date]
	if !ok {
		return
	}
	for i := 0; i < len(c.data[date]); i++ {
		if c.data[date][i].Time == time {
			c.data[date][i] = newEvent
		}
	}

}
