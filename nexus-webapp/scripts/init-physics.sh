#!/bin/bash

# Base directory for the book
BASE_DIR="src/articles/physics"

# Function to create directories, mdx files, and tsx index files
create_chapter() {
  local chapter_path=$1
  local chapter_title=$2
  local sections=("${!3}")
  
  mkdir -p "$chapter_path"
  touch "$chapter_path/index.mdx"
  echo "# $chapter_title" > "$chapter_path/index.mdx"
  
  tsx_content="import { ArticleMeta } from '../../types';\n\n"
  tsx_content+="import content from './index.mdx';\n\n"
  tsx_content+="export const meta: ArticleMeta = {\n"
  tsx_content+="  id: '${chapter_path//\//_}',\n"
  tsx_content+="  title: '$chapter_title',\n"
  tsx_content+="  content,\n"
  tsx_content+="  children: [\n"
  
  for section in "${sections[@]}"; do
    section_path="$chapter_path/$section"
    mkdir -p "$section_path"
    touch "$section_path/index.mdx"
    echo "# $section" > "$section_path/index.mdx"
    
    section_import="${section// /_}"
    tsx_content+="    import { meta as ${section_import,,} } from './$section';\n"
    tsx_content+="    ${section_import,,},\n"
    
    section_tsx_content="import { ArticleMeta } from '../../../types';\n\n"
    section_tsx_content+="import content from './index.mdx';\n\n"
    section_tsx_content+="export const meta: ArticleMeta = {\n"
    section_tsx_content+="  id: '${section_path//\//_}',\n"
    section_tsx_content+="  title: '$section',\n"
    section_tsx_content+="  content,\n"
    section_tsx_content+="  children: []\n"
    section_tsx_content+="};\n"
    
    echo -e "$section_tsx_content" > "$section_path/index.tsx"
  done
  
  tsx_content+="  ]\n"
  tsx_content+="};\n"
  
  echo -e "$tsx_content" > "$chapter_path/index.tsx"
}

# Main book structure
mkdir -p "$BASE_DIR"
touch "$BASE_DIR/index.mdx"
echo "# The Ultimate Interactive Physics Book" > "$BASE_DIR/index.mdx"

base_tsx_content="import { ArticleMeta } from '../types';\n\n"
base_tsx_content+="import content from './index.mdx';\n\n"
base_tsx_content+="export const meta: ArticleMeta = {\n"
base_tsx_content+="  id: '${BASE_DIR//\//_}',\n"
base_tsx_content+="  title: 'The Ultimate Interactive Physics Book',\n"
base_tsx_content+="  content,\n"
base_tsx_content+="  children: [\n"

# Chapter 1: Mechanics
chapter1_sections=(
  "Introduction_to_Mechanics"
  "Newtons_Laws_of_Motion"
  "Kinematics"
  "Dynamics"
  "Work_Energy_and_Power"
  "Conservation_Laws"
)
create_chapter "$BASE_DIR/Chapter_1_Mechanics" "Mechanics" chapter1_sections[@]
base_tsx_content+="    import { meta as chapter_1_mechanics } from './Chapter_1_Mechanics';\n"
base_tsx_content+="    chapter_1_mechanics,\n"

# Chapter 2: Thermodynamics
chapter2_sections=(
  "Introduction_to_Thermodynamics"
  "Temperature_and_Heat"
  "Laws_of_Thermodynamics"
  "Thermodynamic_Processes"
)
create_chapter "$BASE_DIR/Chapter_2_Thermodynamics" "Thermodynamics" chapter2_sections[@]
base_tsx_content+="    import { meta as chapter_2_thermodynamics } from './Chapter_2_Thermodynamics';\n"
base_tsx_content+="    chapter_2_thermodynamics,\n"

# Chapter 3: Electromagnetism
chapter3_sections=(
  "Introduction_to_Electromagnetism"
  "Electrostatics"
  "Electric_Circuits"
  "Magnetism"
  "Electromagnetic_Induction"
)
create_chapter "$BASE_DIR/Chapter_3_Electromagnetism" "Electromagnetism" chapter3_sections[@]
base_tsx_content+="    import { meta as chapter_3_electromagnetism } from './Chapter_3_Electromagnetism';\n"
base_tsx_content+="    chapter_3_electromagnetism,\n"

# Chapter 4: Waves and Optics
chapter4_sections=(
  "Introduction_to_Waves"
  "Sound_Waves"
  "Light_and_Optics"
  "Wave-Particle_Duality"
)
create_chapter "$BASE_DIR/Chapter_4_Waves_and_Optics" "Waves and Optics" chapter4_sections[@]
base_tsx_content+="    import { meta as chapter_4_waves_and_optics } from './Chapter_4_Waves_and_Optics';\n"
base_tsx_content+="    chapter_4_waves_and_optics,\n"

# Chapter 5: Modern Physics
chapter5_sections=(
  "Introduction_to_Modern_Physics"
  "Relativity"
  "Quantum_Mechanics"
  "Particle_Physics"
)
create_chapter "$BASE_DIR/Chapter_5_Modern_Physics" "Modern Physics" chapter5_sections[@]
base_tsx_content+="    import { meta as chapter_5_modern_physics } from './Chapter_5_Modern_Physics';\n"
base_tsx_content+="    chapter_5_modern_physics,\n"

# Chapter 6: Cosmology and Astrophysics
chapter6_sections=(
  "Introduction_to_Cosmology"
  "Big_Bang_Theory"
  "Black_Holes_and_Relativity"
  "Dark_Matter_and_Dark_Energy"
)
create_chapter "$BASE_DIR/Chapter_6_Cosmology_and_Astrophysics" "Cosmology and Astrophysics" chapter6_sections[@]
base_tsx_content+="    import { meta as chapter_6_cosmology_and_astrophysics } from './Chapter_6_Cosmology_and_Astrophysics';\n"
base_tsx_content+="    chapter_6_cosmology_and_astrophysics,\n"

# Chapter 7: Applied Physics
chapter7_sections=(
  "Physics_in_Technology"
  "Physics_in_Medicine"
  "Environmental_Physics"
)
create_chapter "$BASE_DIR/Chapter_7_Applied_Physics" "Applied Physics" chapter7_sections[@]
base_tsx_content+="    import { meta as chapter_7_applied_physics } from './Chapter_7_Applied_Physics';\n"
base_tsx_content+="    chapter_7_applied_physics,\n"

# Chapter 8: Experimental Physics
chapter8_sections=(
  "Introduction_to_Experimental_Methods"
  "Data_Analysis"
  "Significant_Experiments_in_History"
)
create_chapter "$BASE_DIR/Chapter_8_Experimental_Physics" "Experimental Physics" chapter8_sections[@]
base_tsx_content+="    import { meta as chapter_8_experimental_physics } from './Chapter_8_Experimental_Physics';\n"
base_tsx_content+="    chapter_8_experimental_physics,\n"

# Conclusion
conclusion_sections=(
  "Summary_of_Key_Concepts"
  "Future_of_Physics"
  "Further_Reading_and_Resources"
)
create_chapter "$BASE_DIR/Conclusion" "Conclusion" conclusion_sections[@]
base_tsx_content+="    import { meta as conclusion } from './Conclusion';\n"
base_tsx_content+="    conclusion,\n"

# Appendices
appendices_sections=(
  "Mathematical_Tools_for_Physics"
  "Glossary_of_Terms"
)
create_chapter "$BASE_DIR/Appendices" "Appendices" appendices_sections[@]
base_tsx_content+="    import { meta as appendices } from './Appendices';\n"
base_tsx_content+="    appendices,\n"

# Finalize the main index.tsx file
base_tsx_content+="  ]\n"
base_tsx_content+="};\n"

echo -e "$base_tsx_content" > "$BASE_DIR/index.tsx"

echo "Directory structure, MDX files, and TSX metadata files created successfully."
