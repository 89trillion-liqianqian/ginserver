from locust import HttpUser, between, task


class WebsiteUser(HttpUser):
    wait_time = between(5, 15)



    @task
    def keyword(self):
        self.client.get("/getArmyByRarity?rarity=1&unlockArena=1&cvc=1000")
        self.client.get("/getRarity?armyId=10101")
        self.client.get("/getCombatPoints?armyId=10101")
        self.client.get("/getArmyByCvc?cvc=1000")
        self.client.get("/getArmyGroupUnlockArena")