package randomtext

import "math/rand"

func FirstName() string {
	var firstNames = []string{"Natalie", "Aiden", "Brooklyn", "Sebastian", "Lily", "Jayden", "Zoe", "Caleb", "Lucy", "Ryan", "Nora", "Luke", "Ellie", "Andrew", "Stella", "Julian", "Savannah", "Christopher", "Claire", "Nathan", "Violet", "Zachary", "Maya", "Isaac", "Paisley", "John", "Ariana", "Levi", "Hannah", "Anthony", "Hailey", "Leo", "Madeline", "Lincoln", "Aurora", "Jonathan", "Addison", "Dylan", "Skylar", "Carter", "Leah", "Caleb", "Audrey", "Isaiah", "Aria", "Zephyr", "Nova", "Caspian", "Evangeline", "Kairos", "Serenity", "Phoenix", "Amaryllis", "Orion", "Celestia", "Lysandra", "Apollo", "Galatea", "Atlas", "Aurelia", "Cairo", "Delphine", "Soren", "Persephone", "Dante", "Calliope", "Jupiter", "Isolde", "Cyril", "Thalassa", "Icarus", "Elara", "Osiris", "Luna", "Sylvan", "Cressida", "Zephyrine", "Thorne", "Elysia", "Helios", "Vespera", "Calypso", "Caelum", "Aurora", "Ezra", "Oceane", "Cassian", "Astraea", "Maverick", "Zariah", "Indigo", "Jorah", "Nova", "Cassiopeia", "Thorne", "Zara", "Caius", "Elara", "Rune", "Lilith", "Cosmo", "Athena", "Lazarus", "Elowen", "Quill", "Lyra", "Nikolai", "Luna", "Sylvan", "Azura", "Dante", "Selene", "Cyrus", "Astrid", "Orpheus", "Valkyrie", "Kieran", "Aria", "Zephyrine", "Artemis", "Cassian", "Ophelia", "Hawthorne", "Seraphina", "Zenith", "Vesper", "Calista", "Evander", "Eira", "Zephyr", "Thalia", "Caelum", "Seraphine", "Silas", "Evelina", "Orson", "Aurelia", "Odin", "Lavinia", "Peregrine", "Selah", "Aurelius", "Calypso", "Knox", "Eliora", "Zephyra", "Cassiel", "Elowen", "Dorian", "Rhiannon", "Emrys", "Aurelie", "Alistair", "Soleil", "Kai", "Elara", "Caspian", "Ariadne", "Theron", "Ondine", "Lucian", "Seren", "Ezra", "Maeve", "Cyril", "Aurelia", "Orion", "Elysia", "Ronan", "Thalia", "Ophelia", "Ambrose", "Nova", "Lysander", "Calista", "Caius", "Astrid", "Zephyrine", "Selene", "Lysandra", "Clement", "Vesper"}
	randomIndex := rand.Intn(len(firstNames))
	return firstNames[randomIndex]
}