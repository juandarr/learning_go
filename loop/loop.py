from time import time

t0 = time()
net  = 0
for i in range(1, 10**8+1):
    net += i
t = time() - t0
print(net)
print('Operation took {0} secs'.format(t))
