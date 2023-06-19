% Heat flux function of time which has ot be applied on edge 3
P = 0.0035;                                     % Power (W)
length = 0.0003;                                % Length (m)
f=2;                                            % Frequency (Hz)    
W = 2*pi*f;                                     % Pulsation (rad/s)
t_inc = 0.025;                                  % time increment
t_end = 5;                                      % time end
i=1;                                            % matrix increment
for t = 0:t_inc:t_end
    Power(i) = (P+P*cos(2*W*t))/length;         % Power matrix (W/m)
    time(i)=t;                                  % time matrix
    i=i+1;
end

% Geometry
numberOfPDE = 1;                                % Number of equation
model = createpde(numberOfPDE);                 % Create a PDE with numberOfPDE equation

% Define the dimension of the substrate
width = 0.0006; 
height = 0.0003;
gdm = [3 4 0 width width 0 0 0 height height]'; % Define the geometry
g = decsg(gdm, 'S1', ('S1')');                  % Decompose constructive solid
geometryFromEdges(model,g);

%% Set boundary conditions. 
setInitialConditions(model,Ta);                 % Specify the initial temperature on all nodes
applyBoundaryCondition(model,'dirichlet','edge',1,'u',Ta);
applyBoundaryCondition(model,'neumann','edge',3,'g',Power(1));
