class CreateCharacters < ActiveRecord::Migration[5.2]
  def change
    create_table :characters do |t|
      t.string :name
      t.integer :points
      t.integer :strength
      t.integer :ability
      t.integer :resistance
      t.integer :armor
      t.integer :firePower
      t.integer :hp
      t.integer :mp
      t.integer :exp
      t.integer :gold
      t.string :items, array: true, default: []

      t.timestamps
    end
  end
end
