
class Carro:
    def __init__(self, marca, modelo):
        self.marca = marca
        self.modelo = modelo

    def mostrar(self):
        print(f"Marca:{self.marca}, modelo: {self.modelo}")
    

# criando objeto
        
primeiro_carro = Carro("toyota", "corolla")


primeiro_carro.mostrar()
