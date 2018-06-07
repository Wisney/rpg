class Character < ApplicationRecord
  belongs_to :person
  belongs_to :race
  has_many :history
end
