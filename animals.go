package randomtext

import (
	"math/rand"
)

func Animal() string {
	var animals = []string{"Dog", "Cat", "Bird", "Fish", "Frog", "Snake", "Whale", "Lion", "Tiger", "Mouse", "Bear", "Horse", "Rat", "Goat", "Wolf", "Pig", "Duck", "Swan", "Eagle", "Shark", "Sheep", "Hawk", "Crab", "Bat", "Fox", "Seal", "Crow", "Zebra", "Rabbit", "Otter", "Lynx", "Moose", "Chimp", "Lemur", "Squid", "Guppy", "Mule", "Moth", "Panda", "Koala", "Puma", "Gazelle", "Puma", "Ocelot", "Skunk", "Gibbon", "Loris", "Orca", "Finch", "Pika", "Dingo", "Wasp", "Kiwi", "Boa", "Toucan", "Gecko", "Turtle", "Leopard", "Penguin", "Bobcat", "Shrimp", "Falcon", "Jaguar", "Hyena", "Hyrax", "Lizard", "Oyster", "Parrot", "Sloth", "Tarsier", "Walrus", "Yak", "Badger", "Cheetah", "Jaguar", "Vulture", "Gorilla", "Impala", "Marmot", "Ostrich", "Quokka", "Gerbil", "Coyote", "Beagle", "Civet", "Tarsier", "Whippet", "Raccoon", "Hedgehog", "Elephant", "Platypus", "Chinchilla", "Armadillo", "Hamster", "Cheetah", "Alpaca", "Buffalo", "Tortoise", "Cobra", "Antelope", "Caiman", "Carp", "Dolphin", "Giraffe", "Kangaroo", "Llama", "Mongoose", "Wallaby", "Anteater", "Caracal", "Chipmunk", "Gannet", "Gopher", "Jackal", "Langur", "Macaw", "Marmoset", "Pheasant", "Piranha", "Puffin", "Raven", "Redpoll", "Sparrow", "Viper", "Weasel", "Wolverine"}
	randomIndex := rand.Intn(len(animals))
	return animals[randomIndex]
}
