class CreateRaces < ActiveRecord::Migration[5.2]
  def change
    create_table :races do |t|
      t.string :name
      t.text :description
      t.string :changedAttributes, array: true, default: []

      t.timestamps
    end
  end
end
