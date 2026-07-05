import pygame
import random
from logger import log_event
from circleshape import CircleShape
from constants import *


class Asteroid(CircleShape):
  def __init__(self, x: float, y: float, radius: float) -> None:
      super().__init__(x, y, radius)
    
  def draw(self, screen: pygame.Surface) -> None:
      pygame.draw.circle(screen, "grey", self.position, LINE_WIDTH)
      
  def update(self, dt: float) -> None:
      self.position += self.velocity * dt
      
  def split(self) -> None:
      self.kill()
      if self.radius <= ASTEROID_MIN_RADIUS:
          return
      else:
          log_event("asteroid_split")
          old_radius = self.radius
          angle = random.uniform(20, 50)
          new_asteroid_vector_1 = self.velocity.rotate(angle)
          new_asteroid_vector_2 = self.velocity.rotate(-angle)
          new_radius = old_radius - ASTEROID_MIN_RADIUS
          new_asteroid_1 = Asteroid(self.position.x, self.position.y, new_radius)
          new_asteroid_1.velocity = new_asteroid_vector_1 * 1.2
          new_asteroid_2 = Asteroid(self.position.x, self.position.y, new_radius)
          new_asteroid_2.velocity = new_asteroid_vector_2 * 1.2