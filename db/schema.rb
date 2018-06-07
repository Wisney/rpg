# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2018_06_07_004130) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "advantages", force: :cascade do |t|
    t.string "name"
    t.text "description"
    t.string "changedAttributes", default: [], array: true
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  create_table "advantages_characters", id: false, force: :cascade do |t|
    t.bigint "character_id", null: false
    t.bigint "advantage_id", null: false
    t.index ["advantage_id"], name: "index_advantages_characters_on_advantage_id"
    t.index ["character_id"], name: "index_advantages_characters_on_character_id"
  end

  create_table "characters", force: :cascade do |t|
    t.string "name"
    t.integer "points"
    t.integer "strength"
    t.integer "ability"
    t.integer "resistance"
    t.integer "armor"
    t.integer "firePower"
    t.integer "hp"
    t.integer "mp"
    t.integer "exp"
    t.integer "gold"
    t.string "items", default: [], array: true
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.integer "person_id"
    t.integer "race_id"
    t.index ["person_id"], name: "index_characters_on_person_id"
    t.index ["race_id"], name: "index_characters_on_race_id"
  end

  create_table "characters_disadvantages", id: false, force: :cascade do |t|
    t.bigint "character_id", null: false
    t.bigint "disadvantage_id", null: false
    t.index ["character_id"], name: "index_characters_disadvantages_on_character_id"
    t.index ["disadvantage_id"], name: "index_characters_disadvantages_on_disadvantage_id"
  end

  create_table "characters_expertises", id: false, force: :cascade do |t|
    t.bigint "character_id", null: false
    t.bigint "expertise_id", null: false
    t.index ["character_id"], name: "index_characters_expertises_on_character_id"
    t.index ["expertise_id"], name: "index_characters_expertises_on_expertise_id"
  end

  create_table "characters_spells", id: false, force: :cascade do |t|
    t.bigint "character_id", null: false
    t.bigint "spell_id", null: false
    t.index ["character_id"], name: "index_characters_spells_on_character_id"
    t.index ["spell_id"], name: "index_characters_spells_on_spell_id"
  end

  create_table "disadvantages", force: :cascade do |t|
    t.string "name"
    t.text "description"
    t.string "changedAttributes", default: [], array: true
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  create_table "expertises", force: :cascade do |t|
    t.string "name"
    t.text "description"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  create_table "histories", force: :cascade do |t|
    t.text "text"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.integer "character_id"
    t.index ["character_id"], name: "index_histories_on_character_id"
  end

  create_table "people", force: :cascade do |t|
    t.string "username"
    t.string "email"
    t.string "password"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  create_table "races", force: :cascade do |t|
    t.string "name"
    t.text "description"
    t.string "changedAttributes", default: [], array: true
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  create_table "spells", force: :cascade do |t|
    t.string "name"
    t.text "description"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

end
