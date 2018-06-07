class CreateAdvantages < ActiveRecord::Migration[5.2]
  def change
    create_table :advantages do |t|
      t.string :name
      t.text :description
      t.string :changedAttributes, array: true, default: []

      t.timestamps
    end
  end
end
