class CreateAssociations < ActiveRecord::Migration[5.2]
  def change
    add_column :characters, :person_id, :integer
    add_index :characters, :person_id

    add_column :characters, :race_id, :integer
    add_index :characters, :race_id

    add_column :histories, :character_id, :integer
    add_index :histories, :character_id

    create_join_table :characters, :advantages do |t|
      t.index :character_id
      t.index :advantage_id
    end

    create_join_table :characters, :disadvantages do |t|
      t.index :character_id
      t.index :disadvantage_id
    end

    create_join_table :characters, :spells do |t|
      t.index :character_id
      t.index :spell_id
    end

    create_join_table :characters, :expertises do |t|
      t.index :character_id
      t.index :expertise_id
    end
  end
end
